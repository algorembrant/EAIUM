import { useState } from "react";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter, DialogDescription } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

interface CopyConfigDialogProps {
    isOpen: boolean;
    onClose: () => void;
    traderAddress: string;
}

export function CopyConfigDialog({ isOpen, onClose, traderAddress }: CopyConfigDialogProps) {
    const [fixedSize, setFixedSize] = useState("10");
    const [maxPosition, setMaxPosition] = useState("100");
    const [stopLoss, setStopLoss] = useState("15");
    const [takeProfit, setTakeProfit] = useState("30");
    const [isLoading, setIsLoading] = useState(false);

    const handleStartCopy = async () => {
        setIsLoading(true);
        try {
            const response = await fetch("http://localhost:8080/api/copy/start", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    trader_address: traderAddress,
                    enabled: true,
                    fixed_size: parseFloat(fixedSize),
                    max_position: parseFloat(maxPosition),
                    stop_loss_pct: parseFloat(stopLoss) / 100, // Backend expects decimal? Assuming pct in struct comment
                    take_profit_pct: parseFloat(takeProfit) / 100,
                }),
            });

            if (!response.ok) {
                throw new Error("Failed to start copy");
            }

            const data = await response.json();
            console.log("Copy started:", data);
            onClose();
        } catch (error) {
            console.error("Error starting copy:", error);
            // TODO: Show error toast
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <Dialog open={isOpen} onOpenChange={onClose}>
            <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                    <DialogTitle>Copy Configuration</DialogTitle>
                    <DialogDescription>
                        Configure risk parameters for copying {traderAddress.slice(0, 6)}...{traderAddress.slice(-4)}
                    </DialogDescription>
                </DialogHeader>
                <div className="grid gap-4 py-4">
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="size" className="text-right">
                            Bet Size ($)
                        </Label>
                        <Input id="size" value={fixedSize} onChange={(e) => setFixedSize(e.target.value)} className="col-span-3" />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="max" className="text-right">
                            Max Alloc ($)
                        </Label>
                        <Input id="max" value={maxPosition} onChange={(e) => setMaxPosition(e.target.value)} className="col-span-3" />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="sl" className="text-right">
                            Stop Loss (%)
                        </Label>
                        <Input id="sl" value={stopLoss} onChange={(e) => setStopLoss(e.target.value)} className="col-span-3" />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="tp" className="text-right">
                            Take Profit (%)
                        </Label>
                        <Input id="tp" value={takeProfit} onChange={(e) => setTakeProfit(e.target.value)} className="col-span-3" />
                    </div>
                </div>
                <DialogFooter>
                    <Button variant="outline" onClick={onClose}>Cancel</Button>
                    <Button onClick={handleStartCopy} disabled={isLoading}>
                        {isLoading ? "Starting..." : "Start Copying"}
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    );
}
