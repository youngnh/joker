// This file is generated by generate-std.joke script. Do not edit manually!

package math

import (
  
  . "github.com/candid82/joker/core"
)

var mathNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.math"))

var cos_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    x := ExtractNumber(args, 0)
    res := cos(x)
    return MakeDouble(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var hypot_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 2:
    
    p := ExtractNumber(args, 0)
    q := ExtractNumber(args, 1)
    res := hypot(p, q)
    return MakeDouble(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var sin_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    x := ExtractNumber(args, 0)
    res := sin(x)
    return MakeDouble(res)

  default:
    PanicArity(c)
  }
  return NIL
}


func init() {

mathNamespace.ResetMeta(MakeMeta(nil, "Provides basic constants and mathematical functions.", "1.0"))

mathNamespace.InternVar("cos", cos_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("x"))),
    `Returns the cosine of the radian argument x.`, "1.0"))

mathNamespace.InternVar("hypot", hypot_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("p"), MakeSymbol("q"))),
    `Returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.`, "1.0"))

mathNamespace.InternVar("sin", sin_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("x"))),
    `Returns the sine of the radian argument x.`, "1.0"))

}
