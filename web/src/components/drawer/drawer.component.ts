import { coerceBooleanProperty } from '@angular/cdk/coercion';
import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'drawer',
  template: '<ng-content></ng-content>',
  styleUrl: './drawer.component.scss',
  host: {
    class: 'drawer',
    '[class.open]': 'open'
  },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class DrawerComponent {
  @Input()
  get open() { return this._open; }
  set open(v: string | boolean) {
    if (v === '') v = true;
    this._open = coerceBooleanProperty(v);
    this.openChange.emit(this._open);
  }
  private _open: boolean = false;

  @Output()
  readonly openChange = new EventEmitter<boolean>();
}
