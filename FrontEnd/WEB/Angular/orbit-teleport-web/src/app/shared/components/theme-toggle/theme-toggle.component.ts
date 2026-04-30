import { Component, computed, inject } from '@angular/core';
import { ThemeService } from '../../../core/services/theme.service';

@Component({
  selector: 'app-theme-toggle',
  templateUrl: './theme-toggle.component.html',
  styleUrl: './theme-toggle.component.css'
})
export class ThemeToggleComponent {
  private readonly themeService = inject(ThemeService);
  readonly currentTheme = this.themeService.currentTheme;
  readonly label = computed(() => this.currentTheme() === 'day' ? 'Night Mode' : 'Day Mode');

  toggle(): void {
    this.themeService.toggleTheme();
  }
}
