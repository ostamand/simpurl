export default class AnimateDots {
  constructor(element, text, interval = 500, maxDots = 4) {
    this.element = element;
    this.text = text;
    this.interval = interval;
    this.maxDots = maxDots;
  }

  start() {
    this.current = 0;
    this.tick = setInterval(() => {
      this._tickText();
    }, this.interval);
    return this;
  }

  stop() {
    if (this.tick) {
      clearInterval(this.tick);
    }
    return this;
  }

  _tickText() {
    this.current++;
    if (this.current > this.maxDots) {
      this.current = 0;
    }
    this.element.innerText = this.text;
    for (let i = 0; i < this.current; i++) {
      this.element.innerText += ".";
    }
  }
}
