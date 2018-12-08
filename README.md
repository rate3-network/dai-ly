# dai-ly
Etherless payments with Dai

## Inspiration
Payments in stablecoin like Dai should be intuitive to common users. Abstracting away the need for a secondary
Ether balance for gas greatly simplifies wallet management. Send your payments Dai-ly.

## What it does
1. Allows users to have a Dai-only account, and spend Dai by paying Dai fees instead of Ether.
2. Allows Ether holders to be delegated to pay the gas fees in Ether, and get reimbursed in Dai.

## How we built it
Working implementation of [ERC865](https://github.com/ethereum/EIPs/issues/865) for token holders to pay transfer
transactions in tokens instead of gas, in one transaction.

Backend implementation off chain to manage the delegated transactions, as well as matching transactions to the
right delegates.

Frontend wallet implementation for both Dai-only users and Ether gas paying delegates.

## How it works
1. Dai user signs a transaction to send Dai
2. A third party pays for the transaction gas fees on behalf
3. Dai is sent to the destination
4. Third party gets reimbursed Dai from the same transaction that covers the gas fees + a bonus

## Use cases
1. Onboarding new users to a paying with crypto with Dai
2. A payments only Dai wallet without the need to refill Ether
3. Allow users to passively convert Ether into Dai with a slight bonus
