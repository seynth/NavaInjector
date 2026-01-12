using HarmonyLib;
using SafeExamBrowser.WindowsApi;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Milim.patch
{
    [HarmonyPatch(typeof(NativeMethods), nameof(NativeMethods.EmptyClipboard))]
    public class EmptyClipboard
    {
        static bool Prefix()
        {
            return false;
        }
    }

    [HarmonyPatch(typeof(NativeMethods), nameof(NativeMethods.HideWindow))]
    public class HideWindow
    {
        static bool Prefix()
        {
            return false;
        }
    }
}
