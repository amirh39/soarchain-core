// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUnregisterGuard } from "./types/poa/tx";
import { MsgSelectRandomChallenger } from "./types/poa/tx";
import { MsgChallengeService } from "./types/poa/tx";
import { MsgUnregisterRunner } from "./types/poa/tx";
import { MsgV2VChallenge } from "./types/poa/tx";
import { MsgUnregisterChallenger } from "./types/poa/tx";
import { MsgGenGuard } from "./types/poa/tx";
import { MsgUnregisterClient } from "./types/poa/tx";
import { MsgRunnerChallenge } from "./types/poa/tx";
import { MsgSelectRandomRunner } from "./types/poa/tx";
import { MsgGenClient } from "./types/poa/tx";


const types = [
  ["/soarchain.poa.MsgUnregisterGuard", MsgUnregisterGuard],
  ["/soarchain.poa.MsgSelectRandomChallenger", MsgSelectRandomChallenger],
  ["/soarchain.poa.MsgChallengeService", MsgChallengeService],
  ["/soarchain.poa.MsgUnregisterRunner", MsgUnregisterRunner],
  ["/soarchain.poa.MsgV2VChallenge", MsgV2VChallenge],
  ["/soarchain.poa.MsgUnregisterChallenger", MsgUnregisterChallenger],
  ["/soarchain.poa.MsgGenGuard", MsgGenGuard],
  ["/soarchain.poa.MsgUnregisterClient", MsgUnregisterClient],
  ["/soarchain.poa.MsgRunnerChallenge", MsgRunnerChallenge],
  ["/soarchain.poa.MsgSelectRandomRunner", MsgSelectRandomRunner],
  ["/soarchain.poa.MsgGenClient", MsgGenClient],
  
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
    msgUnregisterGuard: (data: MsgUnregisterGuard): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgUnregisterGuard", value: MsgUnregisterGuard.fromPartial( data ) }),
    msgSelectRandomChallenger: (data: MsgSelectRandomChallenger): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgSelectRandomChallenger", value: MsgSelectRandomChallenger.fromPartial( data ) }),
    msgChallengeService: (data: MsgChallengeService): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgChallengeService", value: MsgChallengeService.fromPartial( data ) }),
    msgUnregisterRunner: (data: MsgUnregisterRunner): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgUnregisterRunner", value: MsgUnregisterRunner.fromPartial( data ) }),
    msgV2VChallenge: (data: MsgV2VChallenge): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgV2VChallenge", value: MsgV2VChallenge.fromPartial( data ) }),
    msgUnregisterChallenger: (data: MsgUnregisterChallenger): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgUnregisterChallenger", value: MsgUnregisterChallenger.fromPartial( data ) }),
    msgGenGuard: (data: MsgGenGuard): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgGenGuard", value: MsgGenGuard.fromPartial( data ) }),
    msgUnregisterClient: (data: MsgUnregisterClient): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgUnregisterClient", value: MsgUnregisterClient.fromPartial( data ) }),
    msgRunnerChallenge: (data: MsgRunnerChallenge): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgRunnerChallenge", value: MsgRunnerChallenge.fromPartial( data ) }),
    msgSelectRandomRunner: (data: MsgSelectRandomRunner): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgSelectRandomRunner", value: MsgSelectRandomRunner.fromPartial( data ) }),
    msgGenClient: (data: MsgGenClient): EncodeObject => ({ typeUrl: "/soarchain.poa.MsgGenClient", value: MsgGenClient.fromPartial( data ) }),
    
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
