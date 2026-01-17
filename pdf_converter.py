import time
import os
import sys
from watchdog.observers.polling import PollingObserver
from watchdog.events import FileSystemEventHandler
import markdown
from xhtml2pdf import pisa

class MDHandler(FileSystemEventHandler):
    def on_created(self, event):
        if event.is_directory:
            return
        if event.src_path.endswith('.md'):
            # Check if file actually exists and has content (size > 0)
            print(f"Detected new file: {event.src_path}")
            self.convert_to_pdf(event.src_path)

    def on_modified(self, event):
        if event.is_directory:
            return
        if event.src_path.endswith('.md'):
            print(f"Detected modification: {event.src_path}")
            self.convert_to_pdf(event.src_path)
            
    def on_moved(self, event):
        if event.is_directory:
            return
        if event.dest_path.endswith('.md'):
            print(f"Detected move: {event.dest_path}")
            self.convert_to_pdf(event.dest_path)

    def convert_to_pdf(self, file_path):
        try:
            # Wait a moment for file write to verify it's not empty
            time.sleep(1)
            
            if not os.path.exists(file_path):
                return

            base_name = os.path.basename(file_path)
            file_name_no_ext = os.path.splitext(base_name)[0]
            
            input_dir = os.path.dirname(os.path.abspath(file_path))
            workspace_dir = os.path.dirname(input_dir)
            
            # Verify we are in the correct workspace before writing to sibling 'output'
            if os.path.basename(workspace_dir) != 'md2pdf_workspace':
                # Attempt to handle if we are running from root
                pass

            output_dir = os.path.join(workspace_dir, 'output')
            
            if not os.path.exists(output_dir):
                os.makedirs(output_dir)
            
            output_path = os.path.join(output_dir, f"{file_name_no_ext}.pdf")
            
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
    current_dir = os.getcwd()
    path = os.path.join(current_dir, 'md2pdf_workspace', 'input')
    
    # Ensure directories exist
    if not os.path.exists(path):
        os.makedirs(path)
    
    # Also ensure output exists
    output_path = os.path.join(current_dir, 'md2pdf_workspace', 'output')
    if not os.path.exists(output_path):
        os.makedirs(output_path)
        
    event_handler = MDHandler()
    # Use PollingObserver for better compatibility
    observer = PollingObserver()
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
