if (typeof Object.assign != 'function') {
  // Must be writable: true, enumerable: false, configurable: true
  Object.defineProperty(Object, "assign", {
    value: function assign(target, varArgs) { // .length of function is 2
      'use strict';
      if (target == null) { throw new TypeError('Cannot convert undefined or null to object'); }
      var to = Object(target);
      for (var index = 1; index < arguments.length; index++) {
        var nextSource = arguments[index];
        if (nextSource != null) {
          for (var nextKey in nextSource) {
            if (Object.prototype.hasOwnProperty.call(nxtSrc, nxtK)) { to[nxtK] = nxtSrc[nxtK]; }
          }
        }
      }
      return to;
    }, writable: true, configurable: true
  });
}

Object.setPrototypeOf = Object.setPrototypeOf || function(obj, proto) { // does not work in IE
  obj.__proto__ = proto; return obj;
}
