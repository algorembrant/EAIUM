import sys
import time
import os
import logging
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
import markdown
from xhtml2pdf import pisa
import threading

# Configuration
INPUT_DIR = os.path.abspath("input")
OUTPUT_DIR = os.path.abspath("output")

# Ensure directories exist
os.makedirs(INPUT_DIR, exist_ok=True)
os.makedirs(OUTPUT_DIR, exist_ok=True)

logging.basicConfig(level=logging.INFO,
                    format='%(asctime)s - %(message)s',
                    datefmt='%Y-%m-%d %H:%M:%S')

def convert_md_to_pdf(input_path):
    """Reads a markdown file and converts it to PDF."""
    try:
        filename = os.path.basename(input_path)
        name, _ = os.path.splitext(filename)
        output_path = os.path.join(OUTPUT_DIR, f"{name}.pdf")

        logging.info(f"Processing: {filename}")

        # 1. Read Markdown
        with open(input_path, 'r', encoding='utf-8') as f:
            md_content = f.read()

        # 2. Convert to HTML with tables extension
        # Using 'markdown.extensions.tables' for table support
        # Using 'markdown.extensions.fenced_code' for code blocks
        html_content = markdown.markdown(md_content, extensions=['tables', 'fenced_code'])

        # 3. Add basic styling for PDF
        # xhtml2pdf supports some CSS. We add borders to tables to make them visible.
        full_html = f"""
        <html>
        <head>
        <style>
            @page {{
                size: A4;
                margin: 2cm;
            }}
            body {{
                font-family: Helvetica, sans-serif;
                font-size: 12pt;
            }}
            table {{
                border: 1px solid black;
                border-collapse: collapse;
                width: 100%;
                margin-bottom: 1em;
            }}
            th, td {{
                border: 1px solid black;
                padding: 5px;
                text-align: left;
            }}
            th {{
                background-color: #f2f2f2;
                font-weight: bold;
            }}
            code {{
                font-family: Courier;
                background-color: #f5f5f5;
            }}
            pre {{
                background-color: #f5f5f5;
                padding: 10px;
                border: 1px solid #ccc;
            }}
        </style>
        </head>
        <body>
        {html_content}
        </body>
        </html>
        """

        # 4. Generate PDF
        with open(output_path, "wb") as pdf_file:
            pisa_status = pisa.CreatePDF(
                full_html,                # the HTML to convert
                dest=pdf_file             # file handle to receive result
            )

        if pisa_status.err:
            logging.error(f"Error converting {filename}: {pisa_status.err}")
        else:
            logging.info(f"Successfully created: {output_path}")

    except Exception as e:
        logging.error(f"Failed to convert {input_path}: {e}")

class MDHandler(FileSystemEventHandler):
    def on_created(self, event):
        if not event.is_directory and event.src_path.lower().endswith('.md'):
            # Waiting a brief moment to ensure file write is complete is sometimes helpful,
            # but usually open/read handles it. To be safe:
            time.sleep(0.5)
            convert_md_to_pdf(event.src_path)
    
    # Also handle modified events in case user edits the file in place?
    # The prompt said "put the md file ... directly parser it", so 'created' is primary.
    # 'modified' might trigger too often during saving. Let's stick to 'created' and 'moved'
    # or just 'created' as per "drop file in". 
    # Actually, if I copy paste a file, it's created.
    # If I edit and save, it's modified. 
    # Let's support on_modified too for better UX, but duplicate events might be an issue.
    # We'll stick to 'created' for the "drop file" workflow to be clean.

    def on_moved(self, event):
        if not event.is_directory and event.dest_path.lower().endswith('.md'):
             time.sleep(0.5)
             convert_md_to_pdf(event.dest_path)

if __name__ == "__main__":
    logging.info("Starting Watchdog for MD -> PDF conversion...")
    logging.info(f"Monitoring: {INPUT_DIR}")
    logging.info(f"Output to: {OUTPUT_DIR}")

    event_handler = MDHandler()
    observer = Observer()
    observer.schedule(event_handler, INPUT_DIR, recursive=False)
    observer.start()

    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()
    observer.join()
