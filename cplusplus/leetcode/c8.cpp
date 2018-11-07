
#include<vector>
#include<iostream>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};


bool isSymmetric(TreeNode* root) {
        if(root == NULL) return true;
        vector<TreeNode *> level;
        level.push_back(root);
        while(level.size() > 0 ){
            int mid = level.size() / 2;
            for(int i = 0 ; i <= mid; i++) {
                if(level[i]->val != level[level.size() - i]->val)
                    return false;
            }
            
            vector<TreeNode *> brk(level);
            level.clear();
            for(int i = 0 ; i < brk.size() ; i++) {
                if(brk[i]->left != NULL) 
                    level.push_back(brk[i]->left);
                if(brk[i]->right != NULL)
                    level.push_back(brk[i]->right);
            }
        }
        return true;
    }

int main() {
	TreeNode * root ;
	isSymmetric(root);
}
