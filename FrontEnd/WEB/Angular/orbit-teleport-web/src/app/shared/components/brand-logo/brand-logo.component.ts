import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-brand-logo',
  templateUrl: './brand-logo.component.html',
  styleUrl: './brand-logo.component.css'
})
export class BrandLogoComponent {
  @Input() text = 'entral Teleport';
  @Input() size: 'sm' | 'md' | 'lg' = 'md';
}
