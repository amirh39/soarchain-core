// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgGenClient } from "./types/poa/tx";
import { MsgGenChallenger } from "./types/poa/tx";
import { MsgChallengeService } from "./types/poa/tx";


const types = [
  ["/soarchain.poa.MsgGenClient", MsgGenClient],
  ["/soarchain.poa.MsgGenChallenger", MsgGenChallenger],
  ["/soarchain.poa.MsgChallengeService", MsgChallengeService],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgGenClient: (data: MsgGenClient): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgGenClient", value: MsgGenClient.fromPartial( data ) }),
    msgGenChallenger: (data: MsgGenChallenger): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgGenChallenger", value: MsgGenChallenger.fromPartial( data ) }),
    msgChallengeService: (data: MsgChallengeService): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgChallengeService", value: MsgChallengeService.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
