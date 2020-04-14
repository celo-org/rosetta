package service

var _ Service = (*lazyService)(nil)
var _ Service = (*delayedService)(nil)
