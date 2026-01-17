import { useEffect, useState } from "react";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

interface Position {
    id: number;
    trader_address: string;
    market_id: string;
    entry_price: number;
    current_price: number;
    size: number;
    pnl: number; // calculated on frontend or backend
    is_open: boolean;
}

export function ActivePositions() {
    const [positions, setPositions] = useState<Position[]>([]);

    useEffect(() => {
        // Mock data or fetch from API
        // In a real app, use polling or WebSocket
        const mockPositions: Position[] = [
            {
                id: 1,
                trader_address: "0x123...abc",
                market_id: "Trump 2024",
                entry_price: 0.45,
                current_price: 0.48,
                size: 100,
                pnl: 6.67,
                is_open: true
            },
            {
                id: 2,
                trader_address: "0x789...xyz",
                market_id: "Biden 2024",
                entry_price: 0.30,
                current_price: 0.28,
                size: 50,
                pnl: -6.67,
                is_open: true
            }
        ];
        setPositions(mockPositions);
    }, []);

    return (
        <Card>
            <CardHeader>
                <CardTitle>Active Positions</CardTitle>
            </CardHeader>
            <CardContent>
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>Market</TableHead>
                            <TableHead>Trader</TableHead>
                            <TableHead>Entry</TableHead>
                            <TableHead>Current</TableHead>
                            <TableHead>Size ($)</TableHead>
                            <TableHead>PnL (%)</TableHead>
                            <TableHead>Status</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {positions.map((pos) => {
                            const pnlPct = ((pos.current_price - pos.entry_price) / pos.entry_price) * 100;
                            return (
                                <TableRow key={pos.id}>
                                    <TableCell className="font-medium">{pos.market_id}</TableCell>
                                    <TableCell>{pos.trader_address}</TableCell>
                                    <TableCell>{pos.entry_price.toFixed(3)}</TableCell>
                                    <TableCell>{pos.current_price.toFixed(3)}</TableCell>
                                    <TableCell>{pos.size}</TableCell>
                                    <TableCell className={pnlPct >= 0 ? "text-green-500" : "text-red-500"}>
                                        {pnlPct.toFixed(2)}%
                                    </TableCell>
                                    <TableCell>
                                        <Badge variant="outline">Open</Badge>
                                    </TableCell>
                                </TableRow>
                            );
                        })}
                        {positions.length === 0 && (
                            <TableRow>
                                <TableCell colSpan={7} className="text-center">No active positions</TableCell>
                            </TableRow>
                        )}
                    </TableBody>
                </Table>
            </CardContent>
        </Card>
    );
}
