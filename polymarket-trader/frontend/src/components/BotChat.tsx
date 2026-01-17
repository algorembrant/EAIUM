import { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle, CardFooter } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";

interface Message {
    id: number;
    role: "user" | "bot";
    text: string;
    timestamp: Date;
}

export function BotChat() {
    const [input, setInput] = useState("");
    const [messages, setMessages] = useState<Message[]>([
        { id: 1, role: "bot", text: "Hello! I am your trading assistant. How can I help you today?", timestamp: new Date() }
    ]);

    const handleSend = () => {
        if (!input.trim()) return;

        const userMsg: Message = {
            id: messages.length + 1,
            role: "user",
            text: input,
            timestamp: new Date()
        };

        setMessages(prev => [...prev, userMsg]);
        setInput("");

        // Simulate response
        setTimeout(() => {
            const botMsg: Message = {
                id: messages.length + 2,
                role: "bot",
                text: "I'm currently in demo mode. I can't execute commands yet, but I'm listening!",
                timestamp: new Date()
            };
            setMessages(prev => [...prev, botMsg]);
        }, 1000);
    };

    return (
        <Card className="h-[400px] flex flex-col">
            <CardHeader className="py-3">
                <CardTitle className="text-md">Bot Interface</CardTitle>
            </CardHeader>
            <CardContent className="flex-1 p-0 overflow-hidden">
                <ScrollArea className="h-full p-4">
                    <div className="space-y-4">
                        {messages.map((msg) => (
                            <div key={msg.id} className={`flex ${msg.role === "user" ? "justify-end" : "justify-start"}`}>
                                <div className={`max-w-[80%] rounded-lg px-3 py-2 text-sm ${msg.role === "user" ? "bg-primary text-primary-foreground" : "bg-muted"
                                    }`}>
                                    {msg.text}
                                </div>
                            </div>
                        ))}
                    </div>
                </ScrollArea>
            </CardContent>
            <CardFooter className="p-3 pt-0">
                <form
                    className="flex w-full gap-2"
                    onSubmit={(e) => { e.preventDefault(); handleSend(); }}
                >
                    <Input
                        placeholder="Type a command..."
                        value={input}
                        onChange={(e) => setInput(e.target.value)}
                    />
                    <Button type="submit" size="sm">Send</Button>
                </form>
            </CardFooter>
        </Card>
    );
}
