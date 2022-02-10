export default class AnimateTitle {
  constructor(header, text, interval = 2.5) {
    this.header = header;
    this.text = text;
    this.interval = interval;

    this._generateTitle();
  }

  /**
   * Get max random distinct numbers
   * @param {Number} max
   * @param {Number} n
   */
  _getRandomInt(max, n = 1) {
    const random = (max) => Math.floor(Math.random() * max);
    const ints = [];
    for (let i = 0; i < n; i++) {
      let found = false;
      while (!found) {
        const x = random(max);
        if (!ints.includes(x)) {
          ints.push(x);
          found = true;
        }
      }
    }
    return ints.sort();
  }

  _generateTitle() {
    for (let i = 0; i < this.text.length; i++) {
      const letter = document.createElement("span");
      letter.classList.add("title-default");
      letter.textContent = this.text[i];
      this.header.insertAdjacentElement("beforeend", letter);
    }
  }

  _tickTitle() {
    const n = this.text.length;
    let nChars = this._getRandomInt(n)[0]; // number of chars to highlight
    const indices = this._getRandomInt(n, nChars);
    this._applyHighlight(indices);
  }

  _applyHighlight(indices) {
    for (let i = 0; i < this.text.length; i++) {
      const highlight = indices.includes(i);
      const letter = this.header.childNodes[i];
      if (!highlight && letter.classList.contains("title-color")) {
        letter.classList.remove("title-color");
      }
      if (highlight && !letter.classList.contains("title-color")) {
        letter.classList.add("title-color");
      }
    }
  }

  start() {
    this.tick = setInterval(() => {
      this._tickTitle();
    }, this.interval * 1000);
    return this;
  }

  stop() {
    if (this.tick) {
      clearInterval(this.tick);
    }
    return this;
  }

  highlightLastN(nChars) {
    const n = this.text.length;
    const lastIdx = Math.max(n - 1 - nChars, 0);
    const indices = [];
    for (let i = n - 1; i > lastIdx; i--) {
      indices.push(i);
    }
    this._applyHighlight(indices);
  }
}
