import { HttpEvent, HttpHandlerFn, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs';

import { environment } from '../environments/environment';

export function httpInterceptor(req: HttpRequest<unknown>, next: HttpHandlerFn): Observable<HttpEvent<unknown>> {
  req = req.clone({
    url: `${environment.api.baseUrl}${req.url}`
  });

  return next(req);
}
