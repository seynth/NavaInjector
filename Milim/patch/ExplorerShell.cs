using HarmonyLib;
using SafeExamBrowser.WindowsApi;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Milim.patch
{
    [HarmonyPatch(typeof(ExplorerShell), nameof(ExplorerShell.HideAllWindows))]
    public class UnhideWindow
    {
        static bool Prefix()
        {
            return false;
        }
    }

    [HarmonyPatch(typeof(ExplorerShell), "KillExplorerShell")]
    public class EdotenseiExplorerShell
    {
        static bool Prefix()
        {
            return false;
        }
    }
}
