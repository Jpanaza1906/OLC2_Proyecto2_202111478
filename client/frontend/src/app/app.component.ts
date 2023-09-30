import { Component, ViewChild, ElementRef } from '@angular/core';
import {FileDialogService} from './file-dialog.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'frontend';
  @ViewChild('tswiftCodeTextArea') tswiftCodeTextArea!: ElementRef<HTMLTextAreaElement>;
  constructor(private fileDialogService: FileDialogService) { }

  onOpenFile() {
    this.fileDialogService.openFile().then((content) => {
      if (content) {
        // Do something with the content
        this.tswiftCodeTextArea.nativeElement.value = content;
      }
    });
  }

  onSaveFile() {
    const content = this.tswiftCodeTextArea.nativeElement.value;
    if (content != "") {
      this.fileDialogService.saveFile(content, 'filename.txt');
    } else {
      console.error('No content to save.');
    }
  }

  clearTextArea() {
    this.tswiftCodeTextArea.nativeElement.value = "";
    //focus on text area
    this.tswiftCodeTextArea.nativeElement.focus();
  }

}
