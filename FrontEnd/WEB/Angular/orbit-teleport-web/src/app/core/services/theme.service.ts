import { DOCUMENT } from '@angular/common';
import { Inject, Injectable, signal } from '@angular/core';

type AppTheme = 'day' | 'night';

@Injectable({
  providedIn: 'root'
})
export class ThemeService {
  private readonly storageKey = 'orbit-teleport-theme';
  readonly currentTheme = signal<AppTheme>('day');

  constructor(@Inject(DOCUMENT) private readonly document: Document) {
    this.initTheme();
  }

  toggleTheme(): void {
    const nextTheme: AppTheme = this.currentTheme() === 'day' ? 'night' : 'day';
    this.setTheme(nextTheme);
  }

  setTheme(theme: AppTheme): void {
    this.currentTheme.set(theme);
    this.document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem(this.storageKey, theme);
  }

  private initTheme(): void {
    const storedTheme = localStorage.getItem(this.storageKey) as AppTheme | null;
    const theme: AppTheme = storedTheme === 'night' ? 'night' : 'day';
    this.setTheme(theme);
  }
}
