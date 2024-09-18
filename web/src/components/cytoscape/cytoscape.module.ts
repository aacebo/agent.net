import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CytoscapeComponent } from './cytoscape.component';

@NgModule({
  declarations: [CytoscapeComponent],
  exports: [CytoscapeComponent],
  imports: [CommonModule]
})
export class CytoscapeModule { }
