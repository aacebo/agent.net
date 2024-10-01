import { ChangeDetectionStrategy, Component, ViewEncapsulation, ContentChild } from '@angular/core';

import { DrawerComponent } from './drawer.component';
import { DrawerContentComponent } from './drawer-content.component';

@Component({
  selector: 'drawer-container',
  templateUrl: './drawer-container.component.html',
  styleUrl: './drawer-container.component.scss',
  host: { class: 'drawer-container' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class DrawerContainerComponent {
  @ContentChild(DrawerComponent)
  readonly drawer?: DrawerComponent;

  @ContentChild(DrawerContentComponent)
  readonly content?: DrawerContentComponent;
}
