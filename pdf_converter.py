import time
import os
import sys
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
import markdown
from xhtml2pdf import pisa

class MDHandler(FileSystemEventHandler):
    def process(self, event):
        if event.is_directory:
            return
        if event.src_path.endswith('.md'):
            # Check if file actually exists and has content (size > 0)
            if os.path.exists(event.src_path) and os.path.getsize(event.src_path) > 0:
                print(f"Detected change in: {event.src_path}")
                self.convert_to_pdf(event.src_path)

    def on_created(self, event):
        self.process(event)

    def on_modified(self, event):
        self.process(event)
        
    def on_moved(self, event):
        if event.dest_path.endswith('.md'):
            print(f"Detected move: {event.dest_path}")
             # Create a mock event for the new path
            class MockEvent:
                is_directory = False
                src_path = event.dest_path
            self.convert_to_pdf(event.dest_path)

    def convert_to_pdf(self, file_path):
        try:
            # wait a brief moment to ensure file is fully written and accessible
            time.sleep(1)
            
            # Debounce: check if we just converted this file recently? 
            # For simplicity, we just overwrite.
            
            base_name = os.path.basename(file_path)
            file_name_no_ext = os.path.splitext(base_name)[0]
            
            input_dir = os.path.dirname(file_path)
            workspace_dir = os.path.dirname(input_dir)
            output_dir = os.path.join(workspace_dir, 'output')
            
            if not os.path.exists(output_dir):
                os.makedirs(output_dir)
            
            output_path = os.path.join(output_dir, f"{file_name_no_ext}.pdf")
            
            # Simple check to avoid loop if we were monitoring output too (we aren't)
            
            print(f"Converting {base_name}...")
            
            with open(file_path, 'r', encoding='utf-8') as f:
                text = f.read()
            
            html = markdown.markdown(text)
            
            css = """
            <style>
                body { font-family: sans-serif; }
                h1 { color: #2c3e50; }
                p { line-height: 1.5; }
                code { background-color: #f4f4f4; padding: 2px 5px; border-radius: 3px; }
                pre { background-color: #f4f4f4; padding: 10px; border-radius: 5px; overflow-x: auto;}
            </style>
            """
            
            full_html = f"<html><head>{css}</head><body>{html}</body></html>"
            
            with open(output_path, "wb") as pdf_file:
                pisa_status = pisa.CreatePDF(full_html, dest=pdf_file)
                
            if pisa_status.err:
                print(f"Error converting {base_name}")
            else:
                print(f"Successfully converted {base_name} to {output_path}")
                
        except Exception as e:
            print(f"Failed to convert {file_path}: {e}")

if __name__ == "__main__":
    path = os.path.join(os.getcwd(), 'pdf_workspace', 'input')
    
    # Ensure input directory exists
    if not os.path.exists(path):
        os.makedirs(path)
        
    event_handler = MDHandler()
    observer = Observer()
    observer.schedule(event_handler, path, recursive=False)
    observer.start()
    
    print(f"Monitoring {path} for .md files...")
    print("Press Ctrl+C to stop.")
    
    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()
    observer.join()
