//
// SFML - Simple and Fast Multimedia Library
// Copyright (C) 2007-2009 Laurent Gomila (laurent.gom@gmail.com)
//
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it freely,
// subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented;
// you must not claim that you wrote the original software.
// If you use this software in a product, an acknowledgment
// in the product documentation would be appreciated but is not required.
//
// 2. Altered source versions must be plainly marked as such,
// and must not be misrepresented as being the original software.
//
// 3. This notice may not be removed or altered from any source distribution.
//
#ifndef SFML_CONTEXT_H
#define SFML_CONTEXT_H
// Headers
#include <SFML/Config.h>
#include <SFML/Window/Types.h>
/// Construct a new context
/// return New context
// sfContext* sfContext_Create();
func (self *Context) Create() sfContext* {
    return C.sfContext_Create(self.Cref)
}
/// Destroy an existing context
/// param Context : Context to destroy
// void sfContext_Destroy(sfContext* Context);
func (self *Context) Destroy() void {
    return C.sfContext_Destroy(self.Cref)
}
/// Activate or deactivate a context
/// param Context : Context to activate or deactivate
/// param Active : sfTrue to activate, sfFalse to deactivate
// void sfContext_SetActive(sfContext* Context, sfBool Active);
func (self *Context) SetActive(active Bool) void {
    return C.sfContext_SetActive(self.Cref, active.Cref)
}
#endif // SFML_CONTEXT_H
