# Consensus-Breaking Change Documentation

## What Does Breaking Consensus Mean?

When we talk about breaking consensus in a blockchain, we mean making a change that is not compatible with the old rules. This change makes some nodes in the network disagree with others about what is considered a valid block or transaction. If some nodes use the new rules and others don't, they can end up having different views of the blockchain, which might split the network.

## Why This Change Breaks Consensus

### Change Description

- **What was changed**: I made a rule that says every transaction needs to have a certain amount of tokens or more to be considered valid.

- **How it was changed**: I added a new rule in the code that checks how many tokens are being sent with each transaction. If the amount is too small, the transaction is not allowed.

### Reason for Breaking Consensus

- **Necessity of the change**: To prevent people from sending lots of tiny transactions that don't really mean anything. This can overwhelm the network, making it slow for everyone. The new rule helps keep the network running smoothly.

- **Impact of the change**: Now, any transaction that doesn't meet the minimum amount will be rejected. This means that all nodes in the network need to update to the new version. If some nodes don't update, they will think these transactions are okay and try to include them in the blockchain. This disagreement will cause a split, with some nodes following one version of the blockchain and others following a different one.
