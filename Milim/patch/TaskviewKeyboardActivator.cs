using HarmonyLib;
using SafeExamBrowser.UserInterface.Shared.Activators;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Milim.patch
{
    [HarmonyPatch(typeof(TaskviewKeyboardActivator), "IsActivation")]
    public class IsActivation
    {
        static bool Prefix(ref bool __result)
        {
            __result = false;
            return false;
        }
    }
}
