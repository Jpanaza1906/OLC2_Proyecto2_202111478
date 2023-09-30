import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class FileDialogService {

  constructor() { }

  openFile(): Promise<string | null> {
    return new Promise((resolve) => {
      const input = document.createElement('input');
      input.type = 'file';
      input.accept = '.txt'; // Specify the accepted file types

      input.addEventListener('change', (event: any) => {
        const file = event.target.files[0];
        if (file) {
          const reader = new FileReader();
          reader.onload = () => {
            resolve(reader.result as string);
          };
          reader.readAsText(file);
        } else {
          resolve(null);
        }
      });

      input.click();
    });
  }

  saveFile(content: string, filename: string) {
    const blob = new Blob([content], { type: 'text/plain' });
    const link = document.createElement('a');
    link.href = window.URL.createObjectURL(blob);
    link.download = filename;
    link.click();
  }
}
