document.body.onload = () => {
  fetch("/devices")
    .then((response) => {
      return response.json();
    })
    .then((json) => {
      console.log(json);
      let container = document.createElement("div");
      container.className = "links-container";
      document.body.appendChild(container);
      for (let device of json) {
        let span = document.createElement("span");
        container.appendChild(span);
        let img = document.createElement("img");
        img.src = "icons/" + device.Hostname + ".png";
        const SIZE = 32;
        img.width = SIZE;
        img.height = SIZE;
        span.appendChild(img);
        let a = document.createElement("a");
        a.href = "http://" + device.Href;
        if (device.Alias) {
          a.innerHTML = device.Alias;
        } else {
          a.innerHTML = device.Hostname;
        }
        span.appendChild(a);
      }
    });
};
