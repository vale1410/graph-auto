#hide. 
#show map/2. 

gr(a,b).
gr(b,a).

node(G,X) :- edge(G,X,_).
node(G,X) :- edge(G,_,X).

1 { map(X,Y) : mapping(X,Y) } 1 :- node(a,X).
:- 2 { map(_,Y)}, node(b,Y).
%:- map(X1,Y), map(X2,Y), X1 != X2. 

:-  edge(A,X1,Y1), map(X1,X2), map(Y1,Y2), not edge(B,X2,Y2), X2 != Y2, X1 != Y1, gr(A,B). 

