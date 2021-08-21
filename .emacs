;;; package  --- Hooker - git hooks manager
;;;
;;; Commentary:
;;
;; This package provide Emacs setup to work on hooker project.
;;;
;;; Code:
;; (setq gofmt-command "goimports")
(add-hook 'before-save-hook 'gofmt-before-save)

(provide 'emacs)
;;; .emacs ends here
