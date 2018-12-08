const DelegatableDai = artifacts.require('./DelegatableDai.sol');

const deployContracts = async (deployer, network, accounts) => {
  await deployer.deploy(DelegatableDai);
};

module.exports = function deploy(deployer, network, accounts) {
    return deployer.then(() => deployContracts(deployer, network, accounts));
};
