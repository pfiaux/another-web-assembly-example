(function() {
  'use strict';

  const updateIntervalMs = 2500;
  const go = new Go();
  const colorContainer = document.getElementById("another-color");

  loadAllAssets('lib.wasm')
    .then(([anotherexample, config]) => {
      go.run(anotherexample.instance);

      initializeJSKeyEventListner();
      initializeJSButtonEventListners();

      parseConfig(config);

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
    document.addEventListener('keydown', forwardKeyEvent2Wasm);
    document.addEventListener('keyup', forwardKeyEvent2Wasm);
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
        handleKeyEvent(key);
      });
    });
  }

  function forwardKeyEvent2Wasm(event) {
    if (event.repeat || event.type != 'keydown') {
      return; //Ignore repeat, and keyup
    }
    handleKeyEvent(event.key);
  }

  function updateColor() {
    const newColor = getColor();
    const colorForCSS = `hsl(${newColor.hue}, 50%, 50%)`;
    console.log("Updating color...", colorForCSS);
    colorContainer.style.background = colorForCSS;
  }
})();