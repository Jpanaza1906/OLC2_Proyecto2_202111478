import { Component, ViewChild, ElementRef } from '@angular/core';
import { FileDialogService } from './file-dialog.service'
import { TextService } from './text.service';
import { Entrada, Salida } from './data-model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'frontend';
  //entrada
  @ViewChild('tswiftCodeTextArea') tswiftCodeTextArea!: ElementRef<HTMLTextAreaElement>;
  //salida
  @ViewChild('tswiftCodeOutput') tswiftCodeOutput!: ElementRef<HTMLTextAreaElement>;
  //errores
  @ViewChild('tswiftErroresOutput') tswiftErroresOutput!: ElementRef<HTMLTextAreaElement>;
  //simbolos
  @ViewChild('tswiftSimbolosOutput') tswiftSimbolosOutput!: ElementRef<HTMLTextAreaElement>;
  //funciones
  @ViewChild('tswiftFuncionesOutput') tswiftFuncionesOutput!: ElementRef<HTMLTextAreaElement>;
  //consola
  @ViewChild('tswiftConsolaOutput') tswiftConsolaOutput!: ElementRef<HTMLTextAreaElement>;


  constructor(
    private fileDialogService: FileDialogService,
    private textService: TextService
  ) { }

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

  // Mandar entrada al backend
  
  onCompileCode() {
    const content = this.tswiftCodeTextArea.nativeElement.value;
    if (content.trim() !== ''){
      const entrada:Entrada = {entrada: content};
      this.textService.compileCode(entrada).subscribe(
        (response: Salida) => {
          this.tswiftCodeOutput.nativeElement.value = response.salida;
          this.tswiftErroresOutput.nativeElement.value = response.errores;
          this.tswiftSimbolosOutput.nativeElement.value = response.simbolos;
          this.tswiftFuncionesOutput.nativeElement.value = response.funciones;
          this.tswiftConsolaOutput.nativeElement.value = response.consola;
        },
        (error) => {
          console.log('Error compiling code ',error);
        }
      );
    }else{
      console.error('No content to compile.');
    }
  }

  // tabs
  onKeyDown(event: KeyboardEvent) {
    // Handle the tab key press
    if (event.key === 'Tab') {
      event.preventDefault(); // Prevent the default behavior

      // Insert a tab at the current caret position
      const textarea = this.tswiftCodeTextArea.nativeElement;
      const start = textarea.selectionStart;
      const end = textarea.selectionEnd;

      // Set textarea value to: text before caret + tab + text after caret
      textarea.value = textarea.value.substring(0, start) + '\t' + textarea.value.substring(end);

      // Put caret at right position (add one for the tab)
      textarea.selectionStart = textarea.selectionEnd = start + 1;
    }
  }


}
