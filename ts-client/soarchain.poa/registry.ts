import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUnregisterGuard } from "./types/poa/tx";
import { MsgGenClient } from "./types/poa/tx";
import { MsgUnregisterChallenger } from "./types/poa/tx";
import { MsgUnregisterRunner } from "./types/poa/tx";
import { MsgUnregisterClient } from "./types/poa/tx";
import { MsgSelectRandomChallenger } from "./types/poa/tx";
import { MsgUpdateGuard } from "./types/poa/tx";
import { MsgGenGuard } from "./types/poa/tx";
import { MsgRunnerChallenge } from "./types/poa/tx";
import { MsgChallengeService } from "./types/poa/tx";
import { MsgSelectRandomRunner } from "./types/poa/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/soarchain.poa.MsgUnregisterGuard", MsgUnregisterGuard],
    ["/soarchain.poa.MsgGenClient", MsgGenClient],
    ["/soarchain.poa.MsgUnregisterChallenger", MsgUnregisterChallenger],
    ["/soarchain.poa.MsgUnregisterRunner", MsgUnregisterRunner],
    ["/soarchain.poa.MsgUnregisterClient", MsgUnregisterClient],
    ["/soarchain.poa.MsgSelectRandomChallenger", MsgSelectRandomChallenger],
    ["/soarchain.poa.MsgUpdateGuard", MsgUpdateGuard],
    ["/soarchain.poa.MsgGenGuard", MsgGenGuard],
    ["/soarchain.poa.MsgRunnerChallenge", MsgRunnerChallenge],
    ["/soarchain.poa.MsgChallengeService", MsgChallengeService],
    ["/soarchain.poa.MsgSelectRandomRunner", MsgSelectRandomRunner],
    
];

export { msgTypes }