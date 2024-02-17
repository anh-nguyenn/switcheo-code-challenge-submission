/*
purpose of js: 
1. Initialization of currency rates and symbols.
2. Functions to handle currency selection from the modal, including updating labels and symbols.
3. Conversion logic that dynamically adjusts based on the selected currencies.
4. Input validation to ensure only numeric values are entered.
*/
document.addEventListener("DOMContentLoaded", function () {
  var rates = {
    ETH: 1645.93,
    BTC: 26002.82,
    BLUR: 0.208,
    USD: 1,
    LUNA: 0.406,
    ATOM: 7.19,
    ZIL: 0.017,
    SWTH: 0.004,
    LSI: 67.697,
    WBTC: 26002.82,
    // ....other coins
  };

  var inputCryptoSymbol = "ETH"; // Default input symbol
  var outputCryptoSymbol = "BTC"; // Default output symbol

  // Updates the conversion values based on input or output
  function updateConversion(source) {
    var inputValue = parseFloat(document.getElementById("crypto-input").value);
    var outputValue = parseFloat(
      document.getElementById("crypto-output").value
    );
    var inputUsdValue = 0;
    var outputUsdValue = 0;

    if (source === "input") {
      if (
        !isNaN(inputValue) &&
        rates[inputCryptoSymbol] &&
        rates[outputCryptoSymbol]
      ) {
        inputUsdValue = inputValue * rates[inputCryptoSymbol];
        outputValue =
          (inputValue * rates[inputCryptoSymbol]) / rates[outputCryptoSymbol];
        outputUsdValue = outputValue * rates[outputCryptoSymbol];
        document.getElementById("crypto-output").value = outputValue.toFixed(6);
      } else {
        outputValue = 0;
        outputUsdValue = 0;
        document.getElementById("crypto-output").value = "";
      }
    } else if (source === "output") {
      if (
        !isNaN(outputValue) &&
        rates[inputCryptoSymbol] &&
        rates[outputCryptoSymbol]
      ) {
        inputValue =
          (outputValue * rates[outputCryptoSymbol]) / rates[inputCryptoSymbol];
        inputUsdValue = inputValue * rates[inputCryptoSymbol];
        outputUsdValue = outputValue * rates[outputCryptoSymbol];
        document.getElementById("crypto-input").value = inputValue.toFixed(6);
      } else {
        inputValue = 0;
        inputUsdValue = 0;
        document.getElementById("crypto-input").value = "";
      }
    }

    document.getElementById("crypto-input-usd-value").innerText = isNaN(
      inputUsdValue
    )
      ? "0"
      : inputUsdValue.toFixed(2);
    document.getElementById("crypto-output-usd-value").innerText = isNaN(
      outputUsdValue
    )
      ? "0"
      : outputUsdValue.toFixed(2);
  }

  // Enforce numeric input in the input and output fields
  var enforceNumericInput = function (event) {
    if (
      /[0-9]|\./.test(event.key) ||
      event.keyCode === 8 ||
      event.keyCode === 9 ||
      event.keyCode === 46
    ) {
      return true;
    }
    event.preventDefault();
    return false;
  };

  document
    .getElementById("crypto-input")
    .addEventListener("keydown", enforceNumericInput);
  document
    .getElementById("crypto-output")
    .addEventListener("keydown", enforceNumericInput);

  // Event listeners for input and output fields to trigger conversion
  document
    .getElementById("crypto-input")
    .addEventListener("input", function () {
      updateConversion("input");
    });

  document
    .getElementById("crypto-output")
    .addEventListener("input", function () {
      updateConversion("output");
    });

  // Update the label and symbol when a new currency is selected from the modal
  function updateLabelAndSymbol(type, cryptoName) {
    var labelId =
      type === "pay"
        ? "pay-crypto-select-label"
        : "receive-crypto-select-label";
    var label = document.getElementById(labelId);
    var iconSrc = rates[cryptoName] ? "./images/" + cryptoName + ".svg" : "";
    label.innerHTML = `<i><img src="${iconSrc}" alt="${cryptoName}" /></i> ${cryptoName} <span class="caret"></span>`;

    if (type === "pay") {
      inputCryptoSymbol = cryptoName;
    } else {
      outputCryptoSymbol = cryptoName;
    }

    updateConversion(type === "pay" ? "input" : "output");
  }

  // Attach event listeners to modal buttons for currency selection
  document.querySelectorAll(".crypto-select-btn").forEach(function (button) {
    button.addEventListener("click", function () {
      var cryptoName = this.getAttribute("data-crypto-name");
      var activeLabel = sessionStorage.getItem("activeLabel"); // "pay" or "receive"
      updateLabelAndSymbol(activeLabel, cryptoName);
    });
  });
});
