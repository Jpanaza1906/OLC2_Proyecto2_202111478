import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Entrada, Salida } from './data-model';

@Injectable({
  providedIn: 'root'
})
export class TextService {
  private serveUrl = 'http://localhost:8080'
  constructor(private http: HttpClient) { }

  compileCode(entrada:Entrada): Observable<Salida>{
    return this.http.post<Salida>(`${this.serveUrl}/compilar`, entrada)
  }  
}
