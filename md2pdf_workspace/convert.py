import sys
import time
import os
import logging
from pathlib import Path
import markdown
from weasyprint import HTML, CSS
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler

# Configuration
INPUT_DIR = Path("input")
OUTPUT_DIR = Path("output")

# Ensure directories exist
INPUT_DIR.mkdir(exist_ok=True)
OUTPUT_DIR.mkdir(exist_ok=True)

# Custom CSS for "perfect" rendering
PDF_CSS = """
@page {
    size: A4;
    margin: 2cm;
}
body {
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
    font-size: 10pt;
    line-height: 1.2; /* Single spacing like requested */
    color: #333;
}
h1, h2, h3, h4, h5, h6 {
    color: #000;
    margin-top: 1.5em;
    margin-bottom: 0.5em;
    line-height: 1.2;
}
h1 { font-size: 18pt; border-bottom: 1px solid #ddd; padding-bottom: 0.3em; }
h2 { font-size: 16pt; }
h3 { font-size: 14pt; }
p {
    margin-bottom: 0.8em;
    text-align: justify;
}
code {
    font-family: Consolas, "Courier New", monospace;
    background-color: #f5f5f5;
    padding: 2px 4px;
    border-radius: 3px;
    font-size: 0.9em;
}
pre {
    background-color: #f5f5f5;
    padding: 1em;
    border-radius: 5px;
    overflow-x: auto;
    font-family: Consolas, "Courier New", monospace;
    font-size: 0.9em;
    border: 1px solid #ddd;
}
/* Table Styling */
table {
    border-collapse: collapse;
    width: 100%;
    margin-bottom: 1em;
    font-size: 0.9em;
}
th, td {
    border: 1px solid #ddd;
    padding: 6px 10px; /* Compact padding */
    text-align: left;
}
th {
    background-color: #f2f2f2;
    font-weight: bold;
}
tr:nth-child(even) {
    background-color: #f9f9f9;
}
blockquote {
    border-left: 4px solid #ddd;
    margin: 0;
    padding-left: 1em;
    color: #666;
}
img {
    max-width: 100%;
    height: auto;
    display: block;
    margin: 1em auto;
}
"""

class MDHandler(FileSystemEventHandler):
    def on_created(self, event):
        if event.is_directory:
            return
        
        filepath = Path(event.src_path)
        if filepath.suffix.lower() == ".md":
            self.process_file(filepath)

    def on_modified(self, event):
        # Optional: handle modifications too if user edits the file in place
        if event.is_directory:
            return
        
        filepath = Path(event.src_path)
        if filepath.suffix.lower() == ".md":
            self.process_file(filepath)

    def process_file(self, filepath):
        try:
            print(f"Detected file: {filepath.name}")
            # Wait briefly to ensure file write is complete
            time.sleep(0.5)

            with open(filepath, "r", encoding="utf-8") as f:
                text = f.read()

            # Convert Markdown to HTML
            html_content = markdown.markdown(
                text, 
                extensions=['tables', 'fenced_code', 'nl2br', 'sane_lists']
            )

            # Create destination path
            output_filename = filepath.stem + ".pdf"
            output_path = OUTPUT_DIR / output_filename

            # Wrap in HTML structure
            full_html = f"""
            <!DOCTYPE html>
            <html>
            <head>
                <meta charset="UTF-8">
            </head>
            <body>
                {html_content}
            </body>
            </html>
            """

            # Generate PDF
            print(f"Converting '{filepath.name}' to PDF...")
            HTML(string=full_html, base_url=str(filepath.parent)).write_pdf(
                output_path, stylesheets=[CSS(string=PDF_CSS)]
            )
            print(f"success: PDF saved to '{output_path}'")

        except Exception as e:
            print(f"Error converting file: {e}")

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO,
                        format='%(asctime)s - %(message)s',
                        datefmt='%Y-%m-%d %H:%M:%S')
    
    event_handler = MDHandler()
    observer = Observer()
    observer.schedule(event_handler, str(INPUT_DIR), recursive=False)
    
    print(f"Monitoring '{INPUT_DIR}' for new .md files...")
    print("Press Ctrl+C to stop.")
    
    observer.start()
    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()
    observer.join()
