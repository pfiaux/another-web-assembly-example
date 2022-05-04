(function() {
  'use strict';

  const updateIntervalMs = 2500;
  const go = new Go();
  const colorContainer = document.getElementById("another-color");

  loadAllAssets()
    .then(([anotherexample, config]) => {
      go.run(anotherexample.instance);

      initializeJSKeyEventListner();
      initializeJSButtonEventListners();

      wasmParseConfig(config);

      updateColor();
      setInterval(updateColor, updateIntervalMs);
    });

  function loadAllAssets() {
    return Promise.all([
      loadGoWasm('lib.wasm'),
      loadFileAsBytes(`/config.yaml`),
    ]);
  }

  function loadGoWasm(filename) {
    return fetch(filename)
      .then((wasmLib) => {
        return WebAssembly.instantiateStreaming(wasmLib, go.importObject);
      });
  }

  async function loadFileAsBytes(filename) {
    return fetch(filename)
      .then((response) => response.arrayBuffer())
      .then((buffer) => new Uint8Array(buffer));
  }

  function initializeJSKeyEventListner() {
    document.addEventListener('keyup', (event) => {
      wasmHandleKeyEvent(event.key)
    });
  }

  function initializeJSButtonEventListners() {
    const buttonID2KeyEventMappings = [
      { id: 'key-up', key: 'ArrowUp' },
      { id: 'key-left', key: 'ArrowLeft' },
      { id: 'key-down', key: 'ArrowDown' },
      { id: 'key-right', key: 'ArrowRight'},
    ]
    buttonID2KeyEventMappings.forEach(({id, key}) => {
      const keyButton = document.getElementById(id);
      keyButton.addEventListener('mousedown', (e) => {
        wasmHandleKeyEvent(key);
      });
    });
  }

  function updateColor() {
    const newColor = wasmGetColor();
    const colorForCSS = `hsl(${newColor.hue}, ${newColor.saturation}%, ${newColor.lightness}%)`;
    console.log("Updating color...", colorForCSS);
    colorContainer.style.background = colorForCSS;
  }
})();
