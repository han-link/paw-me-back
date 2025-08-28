import {useNavigate} from "react-router-dom";
import {ComponentType, SVGProps} from "react";
import {ArrowLeftIcon} from '@heroicons/react/24/outline'
import {useHasParent} from '../hooks/useHasParent.ts'

interface HeaderProps {
    action: string | null
    parent: string | null
    icon: ComponentType<SVGProps<SVGSVGElement>> | null
}

export function Header({action, icon: Icon, parent}: HeaderProps) {
    const navigate = useNavigate();
    const hasParent = useHasParent();
    return (
        <header className="sticky top-0 z-10">
            <div className="navbar bg-base-100 shadow-sm flex justify-between">
                <div>
                    {hasParent && (
                        <button onClick={() => parent ? navigate(parent) : navigate('..')}
                                className="btn btn-square btn-ghost">
                            <ArrowLeftIcon/>
                        </button>
                    )}
                </div>
                <div>
                    {action && Icon && (

                        <button onClick={() => navigate(action as string)}
                                className="btn btn-square btn-ghost">
                            <Icon/>
                        </button>
                    )}
                </div>
            </div>
        </header>
    )
}