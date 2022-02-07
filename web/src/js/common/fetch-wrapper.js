export default class FetchWrapper {
  constructor(url) {
    this.url = url;
  }

  async _sendRequest(path, method, body) {
    const options = {
      method: method,
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    };
    if (body) {
      options["body"] = JSON.stringify(body);
    }
    const response = await fetch(this.url + path, options);
    let data = {};
    try {
      data = await response.json();
    } catch (error) {
      console.error(error);
    }
    return [response.status, data];
  }

  post(path, body) {
    return this._sendRequest(path, "POST", body);
  }

  get(path) {
    return this._sendRequest(path, "GET", null);
  }
}
