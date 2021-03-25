import React, { FC, useContext, useEffect, useState } from "react";
import { PokerInfoContext } from "../../pages/games/poker";
import CurrencyInput from "react-currency-input-field";
import { useRouter } from "next/router";
import { useSnackbar } from "notistack";
import { IClass } from "./models/class";

//Selects the class in that the buy in falls. Returns true if a warning should be displayed
const GetClass = (classes, v, setLobby, lobby) => {
    const val = v * 100;
    for (let i = 0; i < classes.length; i++) {
        if (classes[i].min < val && classes[i].max > val) {
            if (i !== lobby.classIndex) {
                setLobby({ class: classes[i], classIndex: i, changeClass: true }); //selected lobby is in a different class.
                return true;
            }
            return false;
        }
    }
    setLobby({ class: {}, classIndex: -1 });
    return v !== 0;
};

export type JoinOptions = {
    buyIn: string;
    lobbyId: string;
    class: string;
};

type JoinProps = {
    onJoin: (options: JoinOptions) => void;
    classes: IClass[];
};

const Join: FC<JoinProps> = ({ onJoin, classes }) => {
    const { lobby, setLobby } = useContext(PokerInfoContext);
    const [buyIn, setBuyIn] = useState<string>("0");
    const { locale } = useRouter();
    const { enqueueSnackbar } = useSnackbar();

    useEffect(() => {
        if (lobby?.class?.min) {
            if (!lobby.changeClass) {
                console.log("Set buyIn: ", lobby.class.min / 100);
                setBuyIn((lobby.class.min / 100).toString());
            } else {
                setLobby({ ...lobby, changeClass: false });
            }
        }
    }, [lobby.class]);

    const onSubmit = (e) => {
        e.preventDefault();

        const values = {
            buyIn: Math.floor(parseFloat(buyIn) * 100).toString(),
            lobbyId: lobby.id,
            class: lobby.classIndex.toString()
        };
        console.log(values);
        if (lobby) onJoin(values);
    };

    if (!classes || !classes.length) {
        return <div>Cant load poker information</div>;
    }
    return (
        <div>
            <form
                onSubmit={onSubmit}
                className="flex justify-center items-center mx-auto my-5 bg-blue-600 w-screen px-1 py-2 rounded shadow-lg"
                style={{ width: "fit-content" }}>
                <CurrencyInput
                    name="buyIn"
                    className="mx-4 p-1 pl-3 rounded outline-none"
                    placeholder={"Buy In Amount"}
                    intlConfig={{ locale: locale, currency: "USD" }}
                    value={buyIn}
                    onValueChange={(val) => {
                        setBuyIn(val);
                    }}
                    onBlur={() => {
                        if (GetClass(classes, parseFloat(buyIn), setLobby, lobby)) {
                            enqueueSnackbar("Entered Buy In was invalid", { variant: "warning" });
                        }
                    }}
                    allowNegativeValue={false}
                />
                <input
                    className="mx-4 p-1 pl-3 rounded outline-none"
                    name="lobbyId"
                    id="lobbyId"
                    placeholder="Lobby Id"
                    type="text"
                    value={lobby?.id ?? ""}
                    onChange={(e) => setLobby({ ...lobby, id: e.target.value, i: -1 })}
                />
                <button
                    className="bg-yellow-500 text-gray-800 hover:bg-yellow-600 transition-colors duration-150 ease-in-out rounded py-1 px-3 mr-3 disabled:opacity-75"
                    type="submit"
                    disabled={!buyIn || lobby.classIndex < 0}>
                    Join
                </button>
            </form>
        </div>
    );
};

export default Join;
