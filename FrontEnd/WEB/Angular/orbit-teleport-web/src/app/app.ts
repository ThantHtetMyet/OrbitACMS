import { Component, computed, inject, signal } from '@angular/core';
import { NavigationEnd, Router, RouterLink, RouterOutlet } from '@angular/router';
import { filter } from 'rxjs';
import { ThemeToggleComponent } from './shared/components/theme-toggle/theme-toggle.component';
import { BrandLogoComponent } from './shared/components/brand-logo/brand-logo.component';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, RouterLink, ThemeToggleComponent, BrandLogoComponent],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  private readonly router = inject(Router);
  private readonly currentUrl = signal(this.router.url);

  protected readonly title = 'entral Teleport';
  protected readonly isSignInPage = computed(() =>
    this.currentUrl().startsWith('/sign-in') || this.currentUrl().startsWith('/login')
  );

  constructor() {
    this.router.events
      .pipe(filter((event) => event instanceof NavigationEnd))
      .subscribe(() => this.currentUrl.set(this.router.url));
  }
}
