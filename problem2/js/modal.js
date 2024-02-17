/*
purpose of js: 
1. update selected field for pop-up modal
2. exchange information of input and output fied after click exchange icon
*/

document.addEventListener("DOMContentLoaded", function () {
  // Keep track of the current selected currencies
  var selectedCurrencies = {
    pay: { name: "ETH", icon: "./images/ETH.svg" }, // default values
    receive: { name: "BTC", icon: "./images/BTC.svg" }, // default values
  };

  // Update label
  function updateLabel(type, cryptoName, iconSrc) {
    var label = document.getElementById(type + "-crypto-select-label");
    if (label) {
      label.innerHTML = `<i><img src="${iconSrc}" alt="${cryptoName}" /></i> ${cryptoName} <span class="caret"></span>`;
      selectedCurrencies[type] = { name: cryptoName, icon: iconSrc };
    }
  }

  document
    .getElementById("pay-crypto-select-label")
    .addEventListener("click", function () {
      sessionStorage.setItem("activeLabel", "pay");
    });

  document
    .getElementById("receive-crypto-select-label")
    .addEventListener("click", function () {
      sessionStorage.setItem("activeLabel", "receive");
    });

  // handle currency selection
  document.querySelectorAll(".crypto-select-btn").forEach(function (button) {
    button.addEventListener("click", function () {
      var cryptoName = this.getAttribute("data-crypto-name");
      var cryptoIconSrc = this.querySelector("img").getAttribute("src");
      var activeLabel = sessionStorage.getItem("activeLabel");
      updateLabel(activeLabel, cryptoName, cryptoIconSrc);
      $("#cryptoModal").modal("hide");
    });
  });

  // event listener for the exchange icon to also swap input and output values
  document
    .querySelector(".fa-exchange-alt")
    .addEventListener("click", function () {
      // Swap the information for both labels
      var tempCurrency = selectedCurrencies["pay"];
      selectedCurrencies["pay"] = selectedCurrencies["receive"];
      selectedCurrencies["receive"] = tempCurrency;

      // Update both labels with the new information
      updateLabel(
        "pay",
        selectedCurrencies["pay"].name,
        selectedCurrencies["pay"].icon
      );
      updateLabel(
        "receive",
        selectedCurrencies["receive"].name,
        selectedCurrencies["receive"].icon
      );

      // Swap the values in the input and output fields
      var tempValue = document.getElementById("crypto-input").value;
      document.getElementById("crypto-input").value =
        document.getElementById("crypto-output").value;
      document.getElementById("crypto-output").value = tempValue;
    });
});
