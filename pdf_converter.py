import time
import os
import sys
import datetime
from watchdog.observers.polling import PollingObserver
from watchdog.events import FileSystemEventHandler
import markdown
from xhtml2pdf import pisa

def log(message, level="INFO"):
    timestamp = datetime.datetime.now().strftime("%H:%M:%S")
    print(f"[{timestamp}] [{level}] {message}")

class MDHandler(FileSystemEventHandler):
    def on_created(self, event):
        if event.is_directory:
            return
        if event.src_path.endswith('.md'):
            log(f"Detected new file: {os.path.basename(event.src_path)}")
            self.convert_to_pdf(event.src_path)

    def on_modified(self, event):
        if event.is_directory:
            return
        if event.src_path.endswith('.md'):
            log(f"Detected modification: {os.path.basename(event.src_path)}")
            self.convert_to_pdf(event.src_path)
            
    def on_moved(self, event):
        if event.is_directory:
            return
        if event.dest_path.endswith('.md'):
            log(f"Detected move: {os.path.basename(event.dest_path)}")
            self.convert_to_pdf(event.dest_path)

    def convert_to_pdf(self, file_path):
        try:
            # Wait a moment for file write
            time.sleep(1)
            
            if not os.path.exists(file_path):
                return

            base_name = os.path.basename(file_path)
            file_name_no_ext = os.path.splitext(base_name)[0]
            
            input_dir = os.path.dirname(os.path.abspath(file_path))
            workspace_dir = os.path.dirname(input_dir)
            
            if os.path.basename(workspace_dir) != 'md2pdf_workspace':
                pass

            output_dir = os.path.join(workspace_dir, 'output')
            
            if not os.path.exists(output_dir):
                os.makedirs(output_dir)
            
            output_path = os.path.join(output_dir, f"{file_name_no_ext}.pdf")
            
            log(f"Starting conversion for {base_name}...")
            
            with open(file_path, 'r', encoding='utf-8') as f:
                text = f.read()
            
            html = markdown.markdown(text, extensions=['tables'])
            
            # CSS Updates:
            # - Table cells: word-wrap to keep content in cells
            # - Line spacing: 1.5 (interpreted from "0.5 line spacing" as 1.5 spacing convention)
            css = """
            <style>
                @page {
                    size: a4 portrait;
                    margin: 2cm;
                }
                body { 
                    font-family: Helvetica, sans-serif; 
                    line-height: 1.5;
                }
                h1 { color: #2c3e50; }
                h2 { color: #34495e; border-bottom: 1px solid #eee; padding-bottom: 5px; }
                p { margin-bottom: 10px; }
                
                code { 
                    background-color: #f4f4f4; 
                    padding: 2px 5px; 
                    border-radius: 3px; 
                    font-family: Courier;
                }
                pre { 
                    background-color: #f4f4f4; 
                    padding: 10px; 
                    border-radius: 5px; 
                    overflow-x: auto;
                    font-family: Courier;
                }
                
                /* Table Styles */
                table {
                    border-collapse: collapse;
                    width: 100%;
                    margin-bottom: 1em;
                    table-layout: fixed; /* helps with column width distribution */
                }
                th, td {
                    border: 1px solid #ddd;
                    padding: 8px;
                    text-align: left;
                    vertical-align: top;
                    word-wrap: break-word; /* Ensure content stays in cell */
                    line-height: 1.5; /* "0.5 line spacing" -> 1.5 value */
                }
                th {
                    background-color: #f8f9fa;
                    font-weight: bold;
                    color: #2c3e50;
                }
                tr:nth-child(even) {
                    background-color: #fcfcfc;
                }
            </style>
            """
            
            full_html = f"<html><head>{css}</head><body>{html}</body></html>"
            
            with open(output_path, "wb") as pdf_file:
                pisa_status = pisa.CreatePDF(full_html, dest=pdf_file)
                
            if pisa_status.err:
                log(f"Error converting {base_name}", "ERROR")
            else:
                log(f"Successfully converted {base_name} -> output/{os.path.basename(output_path)}", "SUCCESS")
                
        except Exception as e:
            log(f"Failed to convert {file_path}: {e}", "CRITICAL")

if __name__ == "__main__":
    current_dir = os.getcwd()
    path = os.path.join(current_dir, 'md2pdf_workspace', 'input')
    
    if not os.path.exists(path):
        os.makedirs(path)
    
    output_path = os.path.join(current_dir, 'md2pdf_workspace', 'output')
    if not os.path.exists(output_path):
        os.makedirs(output_path)
        
    event_handler = MDHandler()
    observer = PollingObserver()
    observer.schedule(event_handler, path, recursive=False)
    observer.start()
    
    log(f"Monitoring directory: {path}")
    log("Press Ctrl+C to stop.")
    
    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()
        log("Stopping observer...")
    observer.join()
