import {render} from "@testing-library/react";
import MessageList, { Message } from "./MessageList";


describe("MessageList", () => {
    const messages: Message[] = [
        { text: "0|Hello" },
        { text: "1|This is an external message" },
        { text: "2|This is an OK message" },
        { text: "3|This is a bad message" },
        { text: "4|This is a pending message" },
    ];

    const okMessages = {
        [messages[2].text]: true,
    };

    const badMessages = {
        [messages[3].text]: true,
    };

    const externalMessages = {
        [messages[1].text]: true,
    };

    it("renders external message correctly", () => {
        const { getByTestId } = render(
            <MessageList
                messages={messages}
                okMessages={okMessages}
                badMessages={badMessages}
                externalMessages={externalMessages}
            />
        );

        const externalMessage = getByTestId("external-message");
        expect(externalMessage).toBeInTheDocument();
    });

    it("renders OK message correctly", () => {
        const { getByTestId } = render(
            <MessageList
                messages={messages}
                okMessages={okMessages}
                badMessages={badMessages}
                externalMessages={externalMessages}
            />
        );

        const okMessage = getByTestId("ok-message");
        expect(okMessage).toBeInTheDocument();
    });

    it("renders bad message correctly", () => {
        const { getByTestId } = render(
            <MessageList
                messages={messages}
                okMessages={okMessages}
                badMessages={badMessages}
                externalMessages={externalMessages}
            />
        );

        const badMessage = getByTestId("bad-message");
        expect(badMessage).toBeInTheDocument();
    });

    it("renders pending message correctly", async () => {
        const { getAllByTestId } = render(
            <MessageList
                messages={messages}
                okMessages={okMessages}
                badMessages={badMessages}
                externalMessages={externalMessages}
            />
        );
        const pendingMessages = getAllByTestId("pending-message");
        expect(pendingMessages.length).toEqual(2);
    });
});
