import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { OverlayModule } from '@angular/cdk/overlay';

import { DrawerComponent } from './drawer.component';
import { DrawerContainerComponent } from './drawer-container.component';
import { DrawerContentComponent } from './drawer-content.component';

@NgModule({
  declarations: [
    DrawerComponent,
    DrawerContainerComponent,
    DrawerContentComponent
  ],
  exports: [
    DrawerComponent,
    DrawerContainerComponent,
    DrawerContentComponent
  ],
  imports: [
    CommonModule,
    OverlayModule
  ]
})
export class DrawerModule { }
