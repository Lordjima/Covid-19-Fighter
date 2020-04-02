import SimpleStorage from "../../truffle/build/contracts/SimpleStorage.json";
import ComplexStorage from "../../truffle/build/contracts/ComplexStorage.json";
import TutorialToken from "../../truffle/build/contracts/TutorialToken.json";

const options = {
  web3: {
    block: false,
    fallback: {
      type: "ws",
      url: "ws://127.0.0.1:9545",
    },
  },
  contracts: [SimpleStorage, ComplexStorage, TutorialToken],
  events: {
    SimpleStorage: ["StorageSet"],
  },
  polls: {
    accounts: 1500,
  },
};

export default options;
