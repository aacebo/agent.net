import { bootstrapApplication } from '@angular/platform-browser';
import 'hammerjs';

import cytoscape from 'cytoscape';
import fcose from 'cytoscape-fcose';

import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';

cytoscape.use(fcose);

bootstrapApplication(AppComponent, appConfig)
  .catch((err) => console.error(err));
