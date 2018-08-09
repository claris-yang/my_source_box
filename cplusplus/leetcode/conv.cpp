

TreeNode* invertTree(TreeNode* root) {
	if(root == NULL) {
		return NULL;
	}

	TreeNode *t = root->left;
	root->left = root->right;
	root->right = t;
        
	invertTree(root->left);
	invertTree(root->right);
	
	return root;
}
