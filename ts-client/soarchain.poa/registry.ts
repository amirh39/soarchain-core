import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgChallengeService } from "./types/poa/tx";
import { MsgSelectRandomRunner } from "./types/poa/tx";
import { MsgGenClient } from "./types/poa/tx";
import { MsgRunnerChallenge } from "./types/poa/tx";
import { MsgUnregisterGuard } from "./types/poa/tx";
import { MsgUnregisterClient } from "./types/poa/tx";
import { MsgUnregisterChallenger } from "./types/poa/tx";
import { MsgGenGuard } from "./types/poa/tx";
import { MsgUpdateGuard } from "./types/poa/tx";
import { MsgSelectRandomChallenger } from "./types/poa/tx";
import { MsgUnregisterRunner } from "./types/poa/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/soarchain.poa.MsgChallengeService", MsgChallengeService],
    ["/soarchain.poa.MsgSelectRandomRunner", MsgSelectRandomRunner],
    ["/soarchain.poa.MsgGenClient", MsgGenClient],
    ["/soarchain.poa.MsgRunnerChallenge", MsgRunnerChallenge],
    ["/soarchain.poa.MsgUnregisterGuard", MsgUnregisterGuard],
    ["/soarchain.poa.MsgUnregisterClient", MsgUnregisterClient],
    ["/soarchain.poa.MsgUnregisterChallenger", MsgUnregisterChallenger],
    ["/soarchain.poa.MsgGenGuard", MsgGenGuard],
    ["/soarchain.poa.MsgUpdateGuard", MsgUpdateGuard],
    ["/soarchain.poa.MsgSelectRandomChallenger", MsgSelectRandomChallenger],
    ["/soarchain.poa.MsgUnregisterRunner", MsgUnregisterRunner],
    
];

export { msgTypes }