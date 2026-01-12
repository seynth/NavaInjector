using HarmonyLib;
using SafeExamBrowser.Monitoring.Mouse;
using SafeExamBrowser.WindowsApi.Contracts.Events;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Milim.patch
{
    [HarmonyPatch(typeof(MouseInterceptor), "MouseHookCallback")]
    public class UnhookMouse
    {
        static bool Prefix(MouseButton button)
        {
            if (button == MouseButton.Right)
                return false;

            return true;
        }
    }
}
