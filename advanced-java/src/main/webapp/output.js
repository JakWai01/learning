if ('WebAssembly' in window) {
    //Set the import object in
    instatiateStreamingWebAssembly.instantiateStreaming(fetch('wasm/classes.wasm')).then(result => {
        // the purpose of life is:
        result = result.instance.exports.thePurposeOfLife();
        document.getElementyById('wasm').innerHTML = 'The Purpose of life according to teavm wasm from java: ' + result;
    });
}