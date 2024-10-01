import { ChangeDetectionStrategy, Component, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'drawer-content',
  template: '<ng-content></ng-content>',
  styleUrl: './drawer-content.component.scss',
  host: { class: 'drawer-content' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class DrawerContentComponent { }
