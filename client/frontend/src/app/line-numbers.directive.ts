import { Directive, ElementRef, HostListener, Renderer2 } from '@angular/core';

@Directive({
  selector: '[appLineNumbers]'
})
export class LineNumbersDirective {
  constructor(private el: ElementRef, private renderer: Renderer2) {}

  @HostListener('input') onInput() {
    this.addLineNumbers();
  }

  @HostListener('scroll') onScroll() {
    this.updateScrollPosition();
  }

  ngAfterViewInit(): void {
    this.addLineNumbers();
    this.updateScrollPosition();
  }

  private addLineNumbers(): void {
    const lines = this.el.nativeElement.value.split('\n').length;
    const lineNumbers = Array.from({ length: lines }, (_, i) => i + 1);

    const lineNumbersElement = this.el.nativeElement.parentElement.querySelector('.line-numbers');
    lineNumbersElement.innerHTML = lineNumbers.join('<br/>');
  }

  private updateScrollPosition(): void {
    const textarea = this.el.nativeElement;
    const lineNumbersElement = this.el.nativeElement.parentElement.querySelector('.line-numbers');
    lineNumbersElement.scrollTop = textarea.scrollTop;
  }
}
