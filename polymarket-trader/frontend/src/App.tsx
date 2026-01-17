import { useState, useEffect } from "react";
import { Copy, TrendingUp, Activity, Wallet, ShieldAlert, Play, Pause } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

interface Trader {
  address: string;
  win_rate: number;
  total_pnl: number;
  trade_count: number;
  is_monitored: boolean;
}

function App() {
  const [traders, setTraders] = useState<Trader[]>([]);
  const [isRunning, setIsRunning] = useState(false);

  // Mock Data Loading
  useEffect(() => {
    // In real app, fetch from API
    setTraders([
      { address: "0x9f8...7b2", win_rate: 0.78, total_pnl: 12500.5, trade_count: 142, is_monitored: false },
      { address: "0xa12...8c9", win_rate: 0.65, total_pnl: 8900.2, trade_count: 98, is_monitored: true },
      { address: "0xb34...1d4", win_rate: 0.92, total_pnl: 45000.0, trade_count: 310, is_monitored: false },
      { address: "0xc56...2e1", win_rate: 0.55, total_pnl: -1200.0, trade_count: 45, is_monitored: false },
    ]);
  }, []);

  return (
    <div className="flex h-screen bg-background text-foreground font-sans selection:bg-primary selection:text-primary-foreground">
      {/* Sidebar */}
      <aside className="w-64 border-r border-border bg-card/50 backdrop-blur-xl hidden md:flex flex-col">
        <div className="p-6">
          <h1 className="text-2xl font-bold bg-gradient-to-r from-blue-400 to-indigo-500 bg-clip-text text-transparent">
            PolyTrader
          </h1>
          <p className="text-xs text-muted-foreground mt-1">High Frequency Copy Bot</p>
        </div>

        <nav className="flex-1 px-4 space-y-2">
          <Button variant="ghost" className="w-full justify-start text-muted-foreground hover:text-foreground">
            <Activity className="mr-2 h-4 w-4" /> Dashboard
          </Button>
          <Button variant="ghost" className="w-full justify-start text-muted-foreground hover:text-foreground">
            <TrendingUp className="mr-2 h-4 w-4" /> Top Traders
          </Button>
          <Button variant="ghost" className="w-full justify-start text-muted-foreground hover:text-foreground">
            <Copy className="mr-2 h-4 w-4" /> Copy Config
          </Button>
          <Button variant="ghost" className="w-full justify-start text-muted-foreground hover:text-foreground">
            <Wallet className="mr-2 h-4 w-4" /> Portfolio
          </Button>
        </nav>

        <div className="p-4 border-t border-border">
          <div className="flex items-center gap-3">
            <div className={`h-2 w-2 rounded-full ${isRunning ? "bg-green-500 animate-pulse" : "bg-red-500"}`} />
            <span className="text-sm font-medium">{isRunning ? "System Active" : "System Paused"}</span>
          </div>
          <Button
            className="w-full mt-3"
            variant={isRunning ? "destructive" : "default"}
            onClick={() => setIsRunning(!isRunning)}
          >
            {isRunning ? <><Pause className="mr-2 h-4 w-4" /> Stop Engine</> : <><Play className="mr-2 h-4 w-4" /> Start Engine</>}
          </Button>
        </div>
      </aside>

      {/* Main Content */}
      <main className="flex-1 flex flex-col overflow-hidden">
        {/* Header */}
        <header className="h-16 border-b border-border flex items-center justify-between px-6 bg-card/50 backdrop-blur-sm">
          <h2 className="text-lg font-semibold">Live Dashboard</h2>
          <div className="flex items-center gap-4">
            <Badge variant="outline" className="text-green-500 border-green-500/50 bg-green-500/10">
              Connected to Polymarket WS
            </Badge>
            <Avatar>
              <AvatarImage src="https://github.com/shadcn.png" />
              <AvatarFallback>CN</AvatarFallback>
            </Avatar>
          </div>
        </header>

        {/* Dashboard Content */}
        <ScrollArea className="flex-1 p-6">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
            <Card className="bg-gradient-to-br from-card to-card/50 border-border/50 shadow-sm">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm font-medium text-muted-foreground">Total PnL</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold text-green-500">+$2,450.00</div>
                <p className="text-xs text-muted-foreground">+18.2% from last month</p>
              </CardContent>
            </Card>
            <Card className="bg-gradient-to-br from-card to-card/50 border-border/50 shadow-sm">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm font-medium text-muted-foreground">Active Positions</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">12</div>
                <p className="text-xs text-muted-foreground">4 copy trades, 8 manual</p>
              </CardContent>
            </Card>
            <Card className="bg-gradient-to-br from-card to-card/50 border-border/50 shadow-sm">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm font-medium text-muted-foreground">Win Rate (24h)</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold text-blue-500">68%</div>
                <p className="text-xs text-muted-foreground">Based on 25 trades</p>
              </CardContent>
            </Card>
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
            {/* Top Traders Leaderboard */}
            <Card className="lg:col-span-2 border-border/50 shadow-md">
              <CardHeader>
                <div className="flex items-center justify-between">
                  <div>
                    <CardTitle>Top Volume Traders</CardTitle>
                    <CardDescription>Real-time ranking based on PnL and consistency.</CardDescription>
                  </div>
                  <Button variant="outline" size="sm">Refresh</Button>
                </div>
              </CardHeader>
              <CardContent>
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Trader</TableHead>
                      <TableHead>Win Rate</TableHead>
                      <TableHead>PnL</TableHead>
                      <TableHead>Status</TableHead>
                      <TableHead className="text-right">Action</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {traders.map((trader) => (
                      <TableRow key={trader.address}>
                        <TableCell className="font-mono text-xs">{trader.address}</TableCell>
                        <TableCell>
                          <span className={trader.win_rate > 0.6 ? "text-green-500 font-medium" : ""}>
                            {(trader.win_rate * 100).toFixed(0)}%
                          </span>
                        </TableCell>
                        <TableCell className={trader.total_pnl >= 0 ? "text-green-500" : "text-red-500"}>
                          ${trader.total_pnl.toLocaleString()}
                        </TableCell>
                        <TableCell>
                          {trader.is_monitored ? <Badge>Active</Badge> : <Badge variant="secondary">Idle</Badge>}
                        </TableCell>
                        <TableCell className="text-right">
                          <Button size="sm" variant={trader.is_monitored ? "destructive" : "default"}>
                            {trader.is_monitored ? "Stop Copy" : "Copy Trade"}
                          </Button>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </CardContent>
            </Card>

            {/* Risk / Activity Feed */}
            <Card className="border-border/50 shadow-md">
              <CardHeader>
                <CardTitle>System Activity</CardTitle>
                <CardDescription>Live events from the engine.</CardDescription>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="flex gap-3 items-start">
                  <div className="h-8 w-8 rounded-full bg-blue-500/10 flex items-center justify-center shrink-0">
                    <Copy className="h-4 w-4 text-blue-500" />
                  </div>
                  <div>
                    <p className="text-sm font-medium">Copied Trade: 0x9f8...</p>
                    <p className="text-xs text-muted-foreground">Bought YES on "Trump vs Biden" @ 0.52</p>
                    <span className="text-[10px] text-muted-foreground">2 mins ago</span>
                  </div>
                </div>
                <div className="flex gap-3 items-start">
                  <div className="h-8 w-8 rounded-full bg-green-500/10 flex items-center justify-center shrink-0">
                    <TrendingUp className="h-4 w-4 text-green-500" />
                  </div>
                  <div>
                    <p className="text-sm font-medium">Take Profit Hit</p>
                    <p className="text-xs text-muted-foreground">Sold 0x9f8... pos for +15%</p>
                    <span className="text-[10px] text-muted-foreground">15 mins ago</span>
                  </div>
                </div>
                <div className="flex gap-3 items-start">
                  <div className="h-8 w-8 rounded-full bg-yellow-500/10 flex items-center justify-center shrink-0">
                    <ShieldAlert className="h-4 w-4 text-yellow-500" />
                  </div>
                  <div>
                    <p className="text-sm font-medium">Risk Check</p>
                    <p className="text-xs text-muted-foreground">Skipped trade: Max exposure reached.</p>
                    <span className="text-[10px] text-muted-foreground">1 hour ago</span>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </ScrollArea>
      </main>
    </div>
  );
}

export default App;
