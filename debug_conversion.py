import markdown
from xhtml2pdf import pisa
import os

file_path = r'pdf_workspace\input\test_doc.md'

try:
    with open(file_path, 'r', encoding='utf-8') as f:
        text = f.read()

    html = markdown.markdown(text)
    css = """<style>body { font-family: sans-serif; }</style>"""
    full_html = f"<html><head>{css}</head><body>{html}</body></html>"

    output_path = r'pdf_workspace\output\debug_test.pdf'
    
    with open(output_path, "wb") as pdf_file:
        pisa_status = pisa.CreatePDF(full_html, dest=pdf_file)

    if pisa_status.err:
        print("Error converting")
    else:
        print(f"Successfully converted to {output_path}")

except Exception as e:
    print(f"Exception: {e}")
