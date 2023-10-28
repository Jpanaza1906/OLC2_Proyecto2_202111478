package compilador

import (
	"strconv"
)

type Conversor struct {
	ctx *Contexto
}

// Constructor for Conversor----------------------------------------------

func NewConversor(ctx *Contexto) *Conversor {
	return &Conversor{
		ctx: ctx,
	}
}

func (c *Conversor) Ampliar(res *Atributos, tipo TipoE) *Atributos {
	switch tipo {
	case Nil:
		c.ctx.AddError("Error: No se puede convertir nil")
		return NewNill()
	case Bool:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("Error: No se puede convertir nil")
			return NewNill()
		case Bool:
			return res
		case Integer:
			c.ctx.AddError("Error: No se puede convertir int a bool")
			return NewNill()
		case Float:
			c.ctx.AddError("Error: No se puede convertir float a bool")
			return NewNill()
		case String:
			c.ctx.AddError("Error: No se puede convertir string a bool")
			return NewNill()
		case Char:
			c.ctx.AddError("Error: No se puede convertir char a bool")
			return NewNill()
		}
	case Integer:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("Error: No se puede convertir nil")
			return NewNill()
		case Bool:
			c.ctx.AddError("Error: No se puede convertir bool a int")
			return NewNill()
		case Integer:
			return res
		case Float:
			//se convierte el string que viene en Dir a un float64
			valorf, err := strconv.ParseFloat(res.Dir, 64)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a int")
				return NewNill()
			}
			respuesta := int(valorf)
			//respuesta se convierte a string
			resp := strconv.Itoa(respuesta)
			return NewInt(resp)
		case String:
			_, err := strconv.Atoi(res.Dir)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a int")
				return NewNill()
			}
			return NewInt(res.Dir)
		case Char:
			_, err := strconv.Atoi(res.Dir)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a int")
				return NewNill()
			}
			return NewInt(res.Dir)
		}
	case Float:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("Error: No se puede convertir nil")
			return NewNill()
		case Bool:
			c.ctx.AddError("Error: No se puede convertir bool a float")
			return NewNill()
		case Integer:
			//convertir la dir a int
			_, err := strconv.Atoi(res.Dir)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a float")
				return NewNill()
			}
			return NewFloat(res.Dir)
		case Float:
			return res
		case String:
			_, err := strconv.ParseFloat(res.Dir, 64)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a float")
				return NewNill()
			}
			respuesta := NewFloat(res.Dir)
			return respuesta
		case Char:
			_, err := strconv.ParseFloat(res.Dir, 64)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a float")
				return NewNill()
			}
			respuesta := NewFloat(res.Dir)
			return respuesta
		}
	case String:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("Error: No se puede convertir nil")
			return NewNill()
		case Bool:
			if res.Dir == "true" {
				return NewString("true")
			} else {
				return NewString("false")
			}
		case Integer:
			//convertir la dir a int
			valori, err := strconv.Atoi(res.Dir)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a string")
				return NewNill()
			}
			respuesta := NewString(strconv.Itoa(valori))
			return respuesta
		case Float:
			//convertir la dir a float
			valorf, err := strconv.ParseFloat(res.Dir, 64)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a string")
				return NewNill()
			}
			respuesta := NewString(strconv.FormatFloat(valorf, 'f', 2, 64))
			return respuesta
		case String:
			return res
		case Char:
			//convertir el num ascci a string
			valori, err := strconv.Atoi(res.Dir)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a string")
				return NewNill()
			}
			respuesta := NewString(strconv.Itoa(valori))
			return respuesta
		}
	}
	c.ctx.AddError("Error: No se puede convertir " + res.Dir + " a " + tipo.String())
	return NewNill()
}
