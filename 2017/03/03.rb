p->(d){def e n,r,c;if r==0 then n**2-1-c elsif c==n-1
then n**2-1-c-r else o n-1,r-1,c end;end;def o n,r,c;
if c==0 then n-1**2+r elsif r==n-1 then(n-1)**2+r+c
else e n-1,r,c-1 end;end;n=Math.sqrt(d).ceil;
(0...n).map{|r|(0...n).map{|c|case(n%2).zero??
e(n,r,c)+1: o(n,r,c)+1 when 1;$x1,$y1=c,r when d;
$x0,$y0=c,r end}};$x0-$x1.abs+$y0-$y1.abs}[ARGV[0].to_i]
