import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUnregisterGuard } from "./types/poa/tx";
import { MsgSelectRandomChallenger } from "./types/poa/tx";
import { MsgClaimRunnerRewards } from "./types/poa/tx";
import { MsgRunnerChallenge } from "./types/poa/tx";
import { MsgChallengeService } from "./types/poa/tx";
import { MsgRegisterFactoryKey } from "./types/poa/tx";
import { MsgSelectRandomRunner } from "./types/poa/tx";
import { MsgUnregisterRunner } from "./types/poa/tx";
import { MsgGenGuard } from "./types/poa/tx";
import { MsgUnregisterChallenger } from "./types/poa/tx";
import { MsgClaimMotusRewards } from "./types/poa/tx";
import { MsgUpdateGuard } from "./types/poa/tx";
import { MsgUnregisterClient } from "./types/poa/tx";
import { MsgGenClient } from "./types/poa/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/soarchain.poa.MsgUnregisterGuard", MsgUnregisterGuard],
    ["/soarchain.poa.MsgSelectRandomChallenger", MsgSelectRandomChallenger],
    ["/soarchain.poa.MsgClaimRunnerRewards", MsgClaimRunnerRewards],
    ["/soarchain.poa.MsgRunnerChallenge", MsgRunnerChallenge],
    ["/soarchain.poa.MsgChallengeService", MsgChallengeService],
    ["/soarchain.poa.MsgRegisterFactoryKey", MsgRegisterFactoryKey],
    ["/soarchain.poa.MsgSelectRandomRunner", MsgSelectRandomRunner],
    ["/soarchain.poa.MsgUnregisterRunner", MsgUnregisterRunner],
    ["/soarchain.poa.MsgGenGuard", MsgGenGuard],
    ["/soarchain.poa.MsgUnregisterChallenger", MsgUnregisterChallenger],
    ["/soarchain.poa.MsgClaimMotusRewards", MsgClaimMotusRewards],
    ["/soarchain.poa.MsgUpdateGuard", MsgUpdateGuard],
    ["/soarchain.poa.MsgUnregisterClient", MsgUnregisterClient],
    ["/soarchain.poa.MsgGenClient", MsgGenClient],
    
];

export { msgTypes }