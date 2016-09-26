#hide. 
#show map/2. 

node(G,X) :- edge(G,X,_).
node(G,X) :- edge(G,_,X).

deg(G,X,N) :- node(G,X), N = [ edge(G,X,_) ].
%deg(G,X,N) :- node(G,X), N = #count{ edge(G,X,Y) : node(G,Y)}.

1 { map(X,Y) : node(2,Y) : deg(2,Y,N) } 1 :- deg(1,X,N).
%1 { map(X,Y) : node(2,Y) , deg(2,Y,N) } 1 :- deg(1,X,N).
:- map(X1,Y), map(X2,Y), X1 != X2. 

:-  edge(1,X1,Y1), map(X1,X2), map(Y1,Y2), not edge(2,X2,Y2), X2 != Y2, X1 != Y1. 
:-  not edge(1,X1,Y1), map(X1,X2), map(Y1,Y2), edge(2,X2,Y2), X2 != Y2, X1 != Y1.

